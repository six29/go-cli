package todo_test

import (
	"os"
	"testing"

	"github.com/six29/go-cli/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q but got %q instead", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q but got %q instead", taskName, l[0].Task)
	}
	if l[0].Done {
		t.Errorf("New task should not be done.")
	}
	l.Complete(1)
	if !l[0].Done {
		t.Errorf("New task should be completed")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}
	tasks := []string{
		"New task 1",
		"New task 2",
		"New task 3",
	}
	for _, v := range tasks {
		l.Add(v)
	}
	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q but got %q", tasks[0], l[0].Task)
	}
	l.Delete(2)
	if len(l) != 2 {
		t.Errorf("Expected %q but got %q", 2, len(l))
	}
	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q but got %q", tasks[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q but got %q", taskName, l1[0].Task)
	}

	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("error creating temporary file: %s", err.Error())
	}

	defer os.Remove(tempFile.Name())

	if errSavingFile := l1.Save(tempFile.Name()); errSavingFile != nil {
		t.Fatalf("Error saving list to file: %s", errSavingFile.Error())
	}
	if errGettingFile := l2.Get(tempFile.Name()); errGettingFile != nil {
		t.Fatalf("Error loading from file: %s", errGettingFile.Error())
	}
	if l1[0].Task != l2[0].Task {
		t.Fatalf("The l1 task %q is not equal to l2 task %q", l1[0].Task, l2[0].Task)
	}

}
