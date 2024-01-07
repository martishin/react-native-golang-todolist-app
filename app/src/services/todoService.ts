import Todo from "../models/Todo"

// replace with server IP address
const baseUrl = "http://192.168.178.150:3000/todos"

export const getTodos = async (): Promise<Todo[]> => {
  const response = await fetch(baseUrl)
  return response.json()
}

export const createTodo = async (todo: Todo): Promise<Todo> => {
  const response = await fetch(baseUrl, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      title: todo.title,
      completed: todo.completed,
    }),
  })
  return response.json()
}

export const updateTodo = async (todo: Todo): Promise<Todo> => {
  const response = await fetch(`${baseUrl}/${todo.id}`, {
    method: "PATCH",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      title: todo.title,
      completed: todo.completed,
    }),
  })
  return response.json()
}

export const deleteTodo = async (id: number): Promise<boolean> => {
  const response = await fetch(`${baseUrl}/${id}`, {
    method: "DELETE",
  })

  return response.status === 204
}
