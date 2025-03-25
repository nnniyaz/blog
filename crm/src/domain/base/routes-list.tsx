import { Dashboard } from '@routes/pages/dashboard'
import { Bio } from '@routes/pages/bio'
import { BioAdd } from '@routes/pages/bio/pages/bio_add.tsx'
import { BioDetails } from '@routes/pages/bio/pages/bio_details.tsx'
import { BioEdit } from '@routes/pages/bio/pages/bio_edit.tsx'
import { Articles } from '@routes/pages/articles'
import { ArticleAdd } from '@routes/pages/articles/pages/article_add.tsx'
import { ArticleDetails } from '@routes/pages/articles/pages/article_details.tsx'
import { ArticleEdit } from '@routes/pages/articles/pages/article_edit.tsx'
import { Projects } from '@routes/pages/projects'
import { ProjectAdd } from '@routes/pages/projects/pages/project_add.tsx'
import { ProjectDetails } from '@routes/pages/projects/pages/project_details.tsx'
import { ProjectEdit } from '@routes/pages/projects/pages/project_edit.tsx'
import { LatestChanges } from '@routes/pages/latest-changes'
import { Users } from '@routes/pages/users'
import { UserAdd } from '@routes/pages/users/pages/user_add.tsx'
import { UserDetails } from '@routes/pages/users/pages/user_details.tsx'
import { UserEdit } from '@routes/pages/users/pages/user_edit.tsx'
import { Profile } from '@routes/pages/profile'
import { Settings } from '@routes/pages/settings'
import { RoutesPaths } from '@domain/base/routes-paths.ts'
import { SignIn } from '@routes/pages/sign-in'

export const PublicRoutesList = [
    // --------------------------------- Login ---------------------------------
    {
        name: 'sign-in',
        path: RoutesPaths.SIGN_IN,
        element: <SignIn />,
    },
]

export const PrivateRoutesList = [
    // ------------------------------- Dashboard -------------------------------
    {
        name: 'dashboard',
        path: RoutesPaths.DASHBOARD,
        element: <Dashboard />,
    },

    // ---------------------------------- Bio ----------------------------------
    {
        name: 'bio',
        path: RoutesPaths.BIO,
        element: <Bio />,
    },
    {
        name: 'bio_add',
        path: RoutesPaths.BIO_ADD,
        element: <BioAdd />,
    },
    {
        name: 'bio_details',
        path: RoutesPaths.BIO_DETAILS,
        element: <BioDetails />,
    },
    {
        name: 'bio_edit',
        path: RoutesPaths.BIO_EDIT,
        element: <BioEdit />,
    },

    // ------------------------------- Articles --------------------------------
    {
        name: 'articles',
        path: RoutesPaths.ARTICLES,
        element: <Articles />,
    },
    {
        name: 'article_add',
        path: RoutesPaths.ARTICLE_ADD,
        element: <ArticleAdd />,
    },
    {
        name: 'article_details',
        path: RoutesPaths.ARTICLE_DETAILS,
        element: <ArticleDetails />,
    },
    {
        name: 'article_edit',
        path: RoutesPaths.ARTICLE_EDIT,
        element: <ArticleEdit />,
    },

    // ------------------------------- Projects --------------------------------
    {
        name: 'projects',
        path: RoutesPaths.PROJECTS,
        element: <Projects />,
    },
    {
        name: 'project_add',
        path: RoutesPaths.PROJECT_ADD,
        element: <ProjectAdd />,
    },
    {
        name: 'project_details',
        path: RoutesPaths.PROJECT_DETAILS,
        element: <ProjectDetails />,
    },
    {
        name: 'project_edit',
        path: RoutesPaths.PROJECT_EDIT,
        element: <ProjectEdit />,
    },

    // ----------------------------- Latest Changes ----------------------------
    {
        name: 'latest_changes',
        path: RoutesPaths.LATEST_CHANGES,
        element: <LatestChanges />,
    },

    // --------------------------------- Users ---------------------------------
    {
        name: 'users',
        path: RoutesPaths.USERS,
        element: <Users />,
    },
    {
        name: 'user_add',
        path: RoutesPaths.USER_ADD,
        element: <UserAdd />,
    },
    {
        name: 'user_details',
        path: RoutesPaths.USER_DETAILS,
        element: <UserDetails />,
    },
    {
        name: 'user_edit',
        path: RoutesPaths.USER_EDIT,
        element: <UserEdit />,
    },

    // -------------------------------- Profile --------------------------------
    {
        name: 'profile',
        path: RoutesPaths.PROFILE,
        element: <Profile />,
    },

    // -------------------------------- Settings -------------------------------
    {
        name: 'settings',
        path: RoutesPaths.SETTINGS,
        element: <Settings />,
    },
]
