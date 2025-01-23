import {MlString} from "../base/mlString/mlString";

export type Project = {
    id: string;
    name: MlString;
    description: MlString;
    coverUri: string;
    appLink: string;
    sourceCodeLink: string;
    isDeleted: boolean;
    createdAt: string;
    updatedAt: string;
}