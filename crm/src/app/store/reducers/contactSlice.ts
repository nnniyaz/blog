import { createSlice } from '@reduxjs/toolkit'

import { Contact } from '@domain/contact/contact.ts'

type State = {
    contacts: Contact[]
    selectedContact: Contact | null
}

const initialState: State = {
    contacts: [],
    selectedContact: null,
}

type SetContactsAction = {
    type: string
    payload: Contact[]
}

type SetSelectedContactAction = {
    type: string
    payload: Contact | null
}

export const contactSlice = createSlice({
    name: 'contact',
    initialState: initialState,
    reducers: {
        setContacts: (state, action: SetContactsAction) => {
            state.contacts = action.payload
        },
        setSelectedContact: (state, action: SetSelectedContactAction) => {
            state.selectedContact = action.payload
        },
    },
})

export default contactSlice.reducer
