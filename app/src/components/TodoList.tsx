import { Divider, List, Spinner, Text, TopNavigation } from "@ui-kitten/components"
import { ReactNode, useEffect, useState } from "react"
import { StyleSheet } from "react-native"

import Todo from "../models/Todo"
import { createTodo, deleteTodo, getTodos, updateTodo } from "../services/todoService"
import SwipeableListItem from "./SwipeableListItem"
import TodoForm from "./TodoForm"

export default function TodoList(): ReactNode {
  const [refreshing, setRefreshing] = useState(false)
  const [todos, setTodos] = useState<Todo[]>([])

  const handleFormSubmit = (todo: Todo) => {
    console.log(`Todo to create: ${JSON.stringify(todo)}`)
    createTodo(todo).then(() => refresh(true))
  }

  const handleDelete = (todo: Todo) => {
    console.log(`Todo to delete: ${JSON.stringify(todo)}`)
    deleteTodo(todo.id!)
      .then(() => refresh(true))
      .catch((error) => console.error("Error deleting todo:", error))
  }

  const handleToggle = (todo: Todo) => {
    console.log(`Todo to toggle: ${JSON.stringify(todo)}`)
    const updatedTodo = { ...todo, completed: !todo.completed }
    updateTodo(updatedTodo)
      .then(() => refresh(true))
      .catch((error) => console.error("Error updating todo:", error))
  }

  const refresh = async (showSpinner = true) => {
    if (showSpinner) {
      setRefreshing(true)
    }
    try {
      const newTodos = await getTodos()
      setTodos(newTodos)
    } catch (error) {
      console.error("Error fetching todos:", error)
    } finally {
      if (showSpinner) {
        setRefreshing(false)
      }
    }
  }

  useEffect(() => {
    refresh(true)

    const interval = setInterval(() => refresh(false), 1000)
    return () => clearInterval(interval)
  }, [])

  const renderSpinner = () => <>{refreshing && <Spinner size="small" />}</>

  return (
    <>
      <TopNavigation title={TitleComponent} accessoryRight={renderSpinner} />
      <Divider />
      <TodoForm handleFormSubmit={handleFormSubmit} />
      <Divider />
      <List
        data={todos}
        ItemSeparatorComponent={Divider}
        renderItem={({ item }) => (
          <SwipeableListItem
            key={item.id}
            todo={item}
            handleDelete={handleDelete}
            handleToggle={handleToggle}
          />
        )}
      />
    </>
  )
}

const TitleComponent = () => (
  <Text category="h1" style={styles.title}>
    Todo App
  </Text>
)

const styles = StyleSheet.create({
  title: {
    fontSize: 20,
  },
})
