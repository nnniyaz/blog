import { createSlice } from '@reduxjs/toolkit'

import { Book } from '@domain/book/book.ts'

type State = {
    books: Book[]
    selectedBook: Book | null
}

const initialState: State = {
    books: [],
    selectedBook: null,
}

type SetBooksAction = {
    type: string
    payload: Book[]
}

type SetSelectedBookAction = {
    type: string
    payload: Book | null
}

export const bookSlice = createSlice({
    name: 'book',
    initialState: initialState,
    reducers: {
        setBooks(state, action: SetBooksAction) {
            state.books = action.payload
        },
        setSelectedBook(state, action: SetSelectedBookAction) {
            state.selectedBook = action.payload
        },
    },
})

export default bookSlice.reducer
