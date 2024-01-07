import { Button, CheckBox, Icon, ListItem } from "@ui-kitten/components"
import { ReactNode } from "react"

import Todo from "../models/Todo"

interface TodoItemProps {
  todo: Todo
  handleDelete: (todo: Todo) => void
  handleToggle: (todo: Todo) => void
}

export default function TodoItem({ todo, handleDelete, handleToggle }: TodoItemProps): ReactNode {
  const DeleteIcon = (props: any) => <Icon {...props} name="trash-2-outline" />

  return (
    <ListItem
      title={todo.title}
      accessoryLeft={() => (
        <CheckBox checked={todo.completed} onChange={() => handleToggle(todo)} />
      )}
      accessoryRight={() => (
        <Button appearance="ghost" accessoryLeft={DeleteIcon} onPress={() => handleDelete(todo)} />
      )}
    />
  )
}
