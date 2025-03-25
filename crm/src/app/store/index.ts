import { configureStore } from '@reduxjs/toolkit'

import authorReducer from './reducers/authorSlice.ts'
import bioReducer from './reducers/bioSlice.ts'
import bookReducer from './reducers/bookSlice.ts'
import contactReducer from './reducers/contactSlice.ts'
import projectReducer from './reducers/projectSlice.ts'
import systemReducer from './reducers/systemSlice.ts'
import userReducer from './reducers/userSlice.ts'
import currentUserReducer from './reducers/currentUserSlice.ts'
import authReducer from './reducers/authSlice.ts'

export const store = configureStore({
    reducer: {
        system: systemReducer,
        author: authorReducer,
        bio: bioReducer,
        book: bookReducer,
        contact: contactReducer,
        project: projectReducer,
        user: userReducer,
        currentUser: currentUserReducer,
        auth: authReducer,
    },
})

export type RootState = ReturnType<typeof store.getState>
