import { Layout } from "@ui-kitten/components"
import { SafeAreaView } from "react-native"

import TodoList from "./TodoList"

export default function HomeScreen() {
  return (
    <Layout style={{ flex: 1, paddingTop: 30 }}>
      <SafeAreaView style={{ flex: 1 }}>
        <TodoList />
      </SafeAreaView>
    </Layout>
  )
}
