import {createSlice} from '@reduxjs/toolkit'

import {Book} from '../../../../domain/book/book'

type State = {
    books: Book[]
    isLoadingGetBooks: boolean
    selectedBook: Book | null
    isLoadingGetSelectedBook: boolean
}

const initialState: State = {
    books: [],
    isLoadingGetBooks: false,
    selectedBook: null,
    isLoadingGetSelectedBook: false
}

type SetBooksAction = {
    type: string
    payload: Book[]
}

type SetIsLoadingGetBooksAction = {
    type: string,
    payload: boolean
}

type SetSelectedBookAction = {
    type: string,
    payload: Book | null
}

type SetIsLoadingGetSelectedBookAction = {
    type: string,
    payload: boolean
}

export const bookSlice = createSlice({
    name: 'book',
    initialState: initialState,
    reducers: {
        setBooks(state, action: SetBooksAction) {
            state.books = action.payload
        },
        setIsLoadingGetBooks: (state, action: SetIsLoadingGetBooksAction) => {
            state.isLoadingGetBooks = action.payload
        },
        setSelectedBook(state, action: SetSelectedBookAction) {
            state.selectedBook = action.payload
        },
        setIsLoadingGetSelectedBook(state, action: SetIsLoadingGetSelectedBookAction) {
            state.isLoadingGetSelectedBook = action.payload
        },
    }
})

export default bookSlice.reducer