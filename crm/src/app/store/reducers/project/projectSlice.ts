import {createSlice} from '@reduxjs/toolkit'

import {Project} from '../../../../domain/project'

type State = {
    projects: Project[],
    isLoadingGetProjects: boolean,
    selectedProject: Project | null,
    isLoadingGetSelectedProject: boolean,
}

const initialState: State = {
    projects: [],
    isLoadingGetProjects: false,
    selectedProject: null,
    isLoadingGetSelectedProject: false
}

type SetProjectsAction = {
    type: string,
    payload: Project[]
}

type SetIsLoadingGetProjectsAction = {
    type: string,
    payload: boolean
}

type SetSelectedProjectAction = {
    type: string,
    payload: Project | null
}

type SetIsLoadingGetSelectedProjectAction = {
    type: string,
    payload: boolean
}

export const projectSlice = createSlice({
    name: 'project',
    initialState: initialState,
    reducers: {
        setProjects: (state, action: SetProjectsAction) => {
            state.projects = action.payload
        },
        setIsLoadingGetProjects: (state, action: SetIsLoadingGetProjectsAction) => {
            state.isLoadingGetProjects = action.payload
        },
        setSelectedProject: (state, action: SetSelectedProjectAction) => {
            state.selectedProject = action.payload
        },
        setIsLoadingGetSelectedProject: (state, action: SetIsLoadingGetSelectedProjectAction) => {
            state.isLoadingGetSelectedProject = action.payload
        }
    }
})

export default projectSlice.reducer