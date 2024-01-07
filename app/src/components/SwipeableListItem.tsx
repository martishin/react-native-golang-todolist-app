import React, { ReactNode, useRef } from "react"
import { Animated, StyleSheet } from "react-native"
import { RectButton, Swipeable } from "react-native-gesture-handler"

import Todo from "../models/Todo"
import TodoItem from "./TodoItem"

interface SwipeableListItemProps {
  todo: Todo
  handleDelete: (todo: Todo) => void
  handleToggle: (todo: Todo) => void
}

export default function SwipeableListItem({
  todo,
  handleDelete,
  handleToggle,
}: SwipeableListItemProps): ReactNode {
  const swipeableRef = useRef(null)

  const handleToggleAndClose = (todo: Todo) => {
    handleToggle(todo)
    swipeableRef.current?.close() // Close the swipeable item
  }

  const renderRightActions = (progress: any, dragX: any) => {
    const trans = dragX.interpolate({
      inputRange: [0, 50, 100, 101],
      outputRange: [-20, 0, 0, 1],
    })

    return (
      <RectButton style={styles.rightAction} onPress={() => handleDelete(todo)}>
        <Animated.Text style={[styles.actionText, { transform: [{ translateX: trans }] }]}>
          Delete
        </Animated.Text>
      </RectButton>
    )
  }

  const renderLeftActions = (progress: any, dragX: any) => {
    const trans = dragX.interpolate({
      inputRange: [-101, -100, -50, 0],
      outputRange: [-1, 0, 0, 20],
    })

    return (
      <RectButton style={styles.leftAction} onPress={() => handleToggle(todo)}>
        <Animated.Text style={{ transform: [{ translateX: trans }] }}>Check</Animated.Text>
      </RectButton>
    )
  }

  return (
    <Swipeable
      ref={swipeableRef}
      renderRightActions={renderRightActions}
      renderLeftActions={renderLeftActions}
      onSwipeableRightOpen={() => handleDelete(todo)}
      onSwipeableLeftOpen={() => handleToggleAndClose(todo)}
    >
      <TodoItem todo={todo} handleDelete={handleDelete} handleToggle={handleToggle} />
    </Swipeable>
  )
}

const styles = StyleSheet.create({
  leftAction: {
    borderRadius: 10,
    flex: 1,
    flexDirection: "column",
    justifyContent: "center",
    marginHorizontal: 16,
    marginVertical: 8,
    overflow: "hidden",
  },
  rightAction: {
    alignItems: "flex-end",
    borderRadius: 10,
    flex: 1,
    flexDirection: "column",
    justifyContent: "center",
    marginHorizontal: 16,
    marginVertical: 8,
    overflow: "hidden",
  },
})
