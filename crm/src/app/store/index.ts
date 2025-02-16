import { configureStore } from '@reduxjs/toolkit'

import authorReducer from './reducers/author/authorSlice.ts'
import bioReducer from './reducers/bio/bioSlice.ts'
import bookReducer from './reducers/book/bookSlice.ts'
import contactReducer from './reducers/contact/contactSlice.ts'
import projectReducer from './reducers/project/projectSlice.ts'
import systemReducer from './reducers/system/systemSlice.ts'
import userReducer from './reducers/user/userSlice.ts'

export const store = configureStore({
    reducer: {
        system: systemReducer,
        author: authorReducer,
        bio: bioReducer,
        book: bookReducer,
        contact: contactReducer,
        project: projectReducer,
        user: userReducer,
    },
})

export type RootState = ReturnType<typeof store.getState>
