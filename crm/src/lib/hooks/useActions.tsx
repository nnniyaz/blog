import { useDispatch } from 'react-redux'
import { bindActionCreators } from 'redux'

import { authorSlice } from '@app/store/reducers/authorSlice.ts'
import { authSlice, signIn, signOut } from '@app/store/reducers/authSlice.ts'
import { bioSlice } from '@app/store/reducers/bioSlice.ts'
import { bookSlice } from '@app/store/reducers/bookSlice.ts'
import { contactSlice } from '@app/store/reducers/contactSlice.ts'
import {
    currentUserSlice,
    getCurrentUser,
} from '@app/store/reducers/currentUserSlice.ts'
import { getProjects, projectSlice } from '@app/store/reducers/projectSlice.ts'
import { systemSlice } from '@app/store/reducers/systemSlice.ts'
import { userSlice } from '@app/store/reducers/userSlice.ts'

export const useActions = () => {
    const dispatch = useDispatch()
    return bindActionCreators(
        {
            ...authorSlice.actions,
            ...authSlice.actions,
            ...bioSlice.actions,
            ...bookSlice.actions,
            ...contactSlice.actions,
            ...currentUserSlice.actions,
            ...projectSlice.actions,
            ...systemSlice.actions,
            ...userSlice.actions,

            // Author Slice

            // Auth Slice
            signIn,
            signOut,

            // Bio Slice

            // Book Slice

            // Contact Slice

            // Current User Slice
            getCurrentUser,

            // Project Slice
            getProjects,

            // System Slice

            // User Slice
        },
        dispatch,
    )
}
