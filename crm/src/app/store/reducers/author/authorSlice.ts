import { createSlice } from '@reduxjs/toolkit'

import { Author } from '@domain/author/author'

type State = {
    author: Author | null
    isLoadingGetAuthor: boolean
}

const initialState: State = {
    author: null,
    isLoadingGetAuthor: false,
}

type SetAuthorAction = {
    type: string
    payload: Author | null
}

type SetIsLoadingGetAuthorAction = {
    type: string
    payload: boolean
}

export const authorSlice = createSlice({
    name: 'author',
    initialState: initialState,
    reducers: {
        setAuthor: (state, action: SetAuthorAction) => {
            state.author = action.payload
        },
        setIsLoadingGetAuthor: (state, action: SetIsLoadingGetAuthorAction) => {
            state.isLoadingGetAuthor = action.payload
        },
    },
})

export default authorSlice.reducer
