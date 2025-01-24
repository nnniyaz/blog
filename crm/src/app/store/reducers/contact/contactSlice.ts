import {createSlice} from '@reduxjs/toolkit'

import {Contact} from '../../../../domain/contact'

type State = {
    contacts: Contact[]
    isLoadingGetContacts: boolean
    selectedContact: Contact | null
    isLoadingGetSelectedContact: boolean
}

const initialState: State = {
    contacts: [],
    isLoadingGetContacts: false,
    selectedContact: null,
    isLoadingGetSelectedContact: false
}

type SetContactsAction = {
    type: string
    payload: Contact[]
}

type SetIsLoadingGetContactsAction = {
    type: string
    payload: boolean
}

type SetSelectedContactAction = {
    type: string
    payload: Contact | null
}

type SetIsLoadingGetSelectedContactAction = {
    type: string
    payload: boolean
}

export const contactSlice = createSlice({
    name: 'contact',
    initialState: initialState,
    reducers: {
        setContacts: (state, action: SetContactsAction) => {
            state.contacts = action.payload
        },
        setIsLoadingGetContacts: (state, action: SetIsLoadingGetContactsAction) => {
            state.isLoadingGetSelectedContact = action.payload
        },
        setSelectedContact: (state, action: SetSelectedContactAction) => {
            state.selectedContact = action.payload
        },
        setIsLoadingSelectedContact: (state, action: SetIsLoadingGetSelectedContactAction) => {
            state.isLoadingGetSelectedContact = action.payload
        }
    }
})

export default contactSlice.reducer
