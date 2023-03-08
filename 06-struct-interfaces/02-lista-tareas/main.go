package main

import "fmt"

// lista de tareas
type TaskList struct {
	tasks []*Task
}

// métodos
func (tl *TaskList) appendTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) removeTask(index int) {
	//obtenemos lo anterior del indice indicado y los que le siguen omitiendo el que se le pasa
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

// Tareas
type Task struct {
	name      string
	desc      string
	completed bool
}

// methods
func (t *Task) toPrint() {
	fmt.Printf("Nombre %s\nDescripción: %s\nCompletado %t\n", t.name, t.desc, t.completed)

}

func (t *Task) markCompleted() {
	t.completed = true
}

func main() {
	t1 := Task{
		name:      "Curso de go",
		desc:      "Completar curso de Go este mes",
		completed: false,
	}
	t2 := Task{
		name:      "Curso de Html",
		desc:      "Completar curso de html esta semana",
		completed: true,
	}

	lista := TaskList{}
	//& se manda la referencia de la memoria y no la copia
	lista.appendTask(&t1)
	lista.appendTask(&t2)

	fmt.Println("Lista de tareas", lista)

	t3 := Task{
		name:      "Curso de css",
		desc:      "Completar curso de css esta semana",
		completed: false,
	}

	lista.appendTask(&t3)
	fmt.Println("Lista de tareas", lista)
	// t1.toPrint()
	// t2.toPrint()
	lista.removeTask(1)

	for i, task := range lista.tasks {
		// task.toPrint()
		fmt.Println(i, task.name)
	}

}
