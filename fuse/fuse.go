package fuse

import (
	"context"
	"os"
	"syscall"
	"log"
	"sync"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	_ "bazil.org/fuse/fs/fstestutil"
)

type FuseServer struct {
	Mountpoint string
	Filesystem FS
	WaitGroup *sync.WaitGroup
	NotifyCallback func(string) error
	receiveInputChans map[string]chan string
}

func (server *FuseServer) Start() {
	defer server.WaitGroup.Done()

	c, err := fuse.Mount(
		server.Mountpoint,
		fuse.FSName("vinfuse"),
		fuse.Subtype("vinfusefs"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	server.Filesystem = FS{
		server: server,
		rootDir: RootDir{
		 fileList: &FileList{
			 files: make([]PuppetFile, 0),
		 },
		},
		maxInode: 1,
	}

	err = fs.Serve(c, server.Filesystem)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *FuseServer) CreateFile(name string) {
	server.Filesystem.addFile(name)

	if server.receiveInputChans == nil {
		server.receiveInputChans = make(map[string]chan string)
	}
	server.receiveInputChans[name] = make(chan string)
}

func (server *FuseServer) SupplyStdin(name string, text string) {
	server.receiveInputChans[name] <- text
}

func (server *FuseServer) DestroyFile(name string) {
	server.Filesystem.removeFile(name)

	delete(server.receiveInputChans, name)
}

type FS struct {
	rootDir  RootDir
	maxInode uint64

	server *FuseServer
}

func (f FS) Root() (fs.Node, error) {
	return f.rootDir, nil
}

func (f *FS) addFile(name string) error {
	f.maxInode += 1
	f.rootDir.fileList.files = append(f.rootDir.fileList.files, PuppetFile{
		Inode: f.maxInode,
	 	Name:  name,
	 	server: f.server,
	})
	return nil
}

func (f *FS) removeFile(name string) error {
	f.maxInode -= 1
	for i, file := range f.rootDir.fileList.files {
		if file.Name == name {
			f.rootDir.fileList.files = append(f.rootDir.fileList.files[:i], f.rootDir.fileList.files[i+1:]...)
			return nil
		}
 }
	return nil
}

type FileList struct {
	files []PuppetFile
}

type RootDir struct {
	fileList *FileList
}

func (RootDir) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = 1
	a.Mode = os.ModeDir | 0o555
	return nil
}

func (rd RootDir) Lookup(ctx context.Context, name string) (fs.Node, error) {
	for _, f := range rd.fileList.files {
	 	if f.Name == name {
			return f, nil
	 	}
	}
	return nil, syscall.ENOENT
}

func (rd RootDir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	dirents := make([]fuse.Dirent, len(rd.fileList.files))

	for i, f := range rd.fileList.files {
		dirents[i] = fuse.Dirent{Inode: f.Inode, Name: f.Name, Type: fuse.DT_File}
	}

	return dirents, nil
}

type PuppetFile struct {
	Inode uint64
	Name  string

	server *FuseServer
}

func (pf PuppetFile) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = pf.Inode
	a.Mode = 0o444
	a.Size = uint64(9999999) // temporary fix
	return nil
}

func (pf PuppetFile) ReadAll(ctx context.Context) ([]byte, error) {
	pf.server.NotifyCallback(pf.Name)

	text := <- pf.server.receiveInputChans[pf.Name]
	return []byte(text), nil
}