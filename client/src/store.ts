import { configureStore } from '@reduxjs/toolkit'

function locationReducer(state = { value: "" }, action: any) {
	switch (action.type) {
		case 'get':
			return state
		case 'update':
			return state
		default:
			return state
	}
}

export const store = configureStore({
	reducer: {
		location: locationReducer,
	},
})

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch
