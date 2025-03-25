import { createSlice } from '@reduxjs/toolkit'

import { Project } from '@domain/project/project.ts'

type State = {
    projects: Project[]
    isLoadingGetProjects: boolean
    selectedProject: Project | null
    isLoadingGetSelectedProject: boolean
}

const initialState: State = {
    projects: [],
    isLoadingGetProjects: false,
    selectedProject: null,
    isLoadingGetSelectedProject: false,
}

type SetProjectsAction = {
    type: string
    payload: Project[]
}

type SetSelectedProjectAction = {
    type: string
    payload: Project | null
}

export const projectSlice = createSlice({
    name: 'project',
    initialState: initialState,
    reducers: {
        setProjects: (state, action: SetProjectsAction) => {
            state.projects = action.payload
        },
        setSelectedProject: (state, action: SetSelectedProjectAction) => {
            state.selectedProject = action.payload
        },
    },
})

export default projectSlice.reducer
