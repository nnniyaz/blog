export enum RoutesPaths {
    // New entity routes must have only one slash in the path
    // Nested routes must have more than one slash in the path

    SIGN_IN = '/sign-in',

    DASHBOARD = '/',

    BIO = '/bio',
    BIO_ADD = '/bio/add',
    BIO_DETAILS = '/bio/:id',
    BIO_EDIT = '/bio/:id/edit',

    ARTICLES = '/articles',
    ARTICLE_ADD = '/articles/add',
    ARTICLE_DETAILS = '/articles/:id',
    ARTICLE_EDIT = '/articles/:id/edit',

    PROJECTS = '/projects',
    PROJECT_ADD = '/projects/add',
    PROJECT_DETAILS = '/projects/:id',
    PROJECT_EDIT = '/projects/:id/edit',

    LATEST_CHANGES = '/latest-changes',

    USERS = '/users',
    USER_ADD = '/users/add',
    USER_DETAILS = '/users/:id',
    USER_EDIT = '/users/:id/edit',

    PROFILE = '/profile',
    SETTINGS = '/settings',
}
