import { createSlice } from '@reduxjs/toolkit'

import { Author } from '@domain/author/author.ts'

type State = {
    author: Author | null
}

const initialState: State = {
    author: null,
}

type SetAuthorAction = {
    type: string
    payload: Author | null
}

export const authorSlice = createSlice({
    name: 'author',
    initialState: initialState,
    reducers: {
        setAuthor: (state, action: SetAuthorAction) => {
            state.author = action.payload
        },
    },
})

export default authorSlice.reducer
