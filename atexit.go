package atexit

var functions []func()

// run me in main()
func AtExit() {
	for _, f := range functions {
		f()
	}
}

// add functions to call on program exit
func Add(y ...func()) {
	functions = append(functions, y...)
}
