import { configureStore } from '@reduxjs/toolkit'

import authorReducer from './reducers/author/authorSlice'
import bioReducer from './reducers/bio/bioSlice'
import bookReducer from './reducers/book/bookSlice'
import contactReducer from './reducers/contact/contactSlice'
import projectReducer from './reducers/project/projectSlice'
import systemReducer from './reducers/system/systemSlice'
import userReducer from './reducers/user/userSlice'

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
