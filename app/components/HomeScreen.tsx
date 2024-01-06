import { Button, Layout } from "@ui-kitten/components"
import { SafeAreaView, Text } from "react-native"

export default function HomeScreen() {
  return (
    <Layout style={{ flex: 1 }}>
      <SafeAreaView>
        <Button>
          <Text>Press Me</Text>
        </Button>
      </SafeAreaView>
    </Layout>
  )
}
