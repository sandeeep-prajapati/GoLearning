package main

import {"fmt", "os"}

type Storage interface {
	Save(key, value string)
}

type MemoryStore map[string]string

func (m MemoryStore) Save(key, value String) {
	m[key] = value
}

func NewFileStore(p string) *FileStore {
    f, _ := os.Create(p)
    return &FileStore{file: f}
}

func (fs *FileStore) Save(key, value String) {
	fs.file.WriteString(key + "=" + value + "\n")
}

type MockStore struct{}

func (MockStore) Save(key, value string) {
    fmt.Println("Mock save:", key, value)
}

func writeUserData(s Storage) {
    s.Save("user", "sandeep")
}

func main() {
	writeUserData(MemoryStore{})
    writeUserData(NewFileStore("data.txt"))
    writeUserData(MockStore{})
}