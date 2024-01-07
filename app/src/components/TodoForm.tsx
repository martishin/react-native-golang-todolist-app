import { Button, Input, Layout, Text } from "@ui-kitten/components"
import { useState } from "react"
import { Keyboard, StyleSheet, View } from "react-native"

import Todo from "../models/Todo"

interface TodoFormProps {
  handleFormSubmit: (todo: Todo) => void
}

export default function TodoForm({ handleFormSubmit }: TodoFormProps) {
  const [title, setTitle] = useState("")

  const handleSubmit = () => {
    handleFormSubmit({ title, completed: false })
    setTitle("")
    Keyboard.dismiss()
  }

  return (
    <Layout style={styles.rowContainer} level="1">
      <Input
        style={styles.input}
        status="basic"
        placeholder="New Todo"
        value={title}
        onChangeText={(nextTitle) => setTitle(nextTitle)}
        onSubmitEditing={handleSubmit}
      />
      <View style={styles.controlContainer}>
        <Button size="tiny" onPress={handleSubmit}>
          <Text>Create</Text>
        </Button>
      </View>
    </Layout>
  )
}

const styles = StyleSheet.create({
  controlContainer: {
    backgroundColor: "#3366FF",
    borderRadius: 4,
    margin: 2,
    padding: 6,
  },
  input: {
    flex: 1,
    margin: 2,
  },
  rowContainer: {
    alignItems: "center",
    flexDirection: "row",
    justifyContent: "space-between",
  },
})
