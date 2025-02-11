import { Db, MongoClient } from 'mongodb'

let MONGO_URI: string | null = null

let cachedClient: MongoClient | null = null
let articleDb: Db | null = null
let authorDb: Db | null = null
let bioDb: Db | null = null
let bookDb: Db | null = null
let contactDb: Db | null = null
let projectDb: Db | null = null
let sessionDb: Db | null = null
let userDb: Db | null = null

interface connectToDatabaseReturnProps {
    client: MongoClient
}

export async function connectToDatabase(): Promise<connectToDatabaseReturnProps> {
    if (!MONGO_URI) {
        MONGO_URI = process.env.MONGO_URI || ''
    }

    // check the MongoDB URI
    if (!MONGO_URI) {
        throw new Error('Define the MONGO_URI environmental variable')
    }

    // check the cached.
    if (cachedClient) {
        // load from cache
        return {
            client: cachedClient,
        }
    }

    // Connect to cluster
    const client: MongoClient = new MongoClient(MONGO_URI)
    await client.connect()

    // set cache
    cachedClient = client

    return {
        client: cachedClient,
    }
}

export async function getArticleDb(): Promise<Db> {
    if (articleDb) {
        return articleDb
    }

    const { client } = await connectToDatabase()
    articleDb = client.db('article')
    return articleDb
}

export async function getAuthorDb(): Promise<Db> {
    if (authorDb) {
        return authorDb
    }

    const { client } = await connectToDatabase()
    authorDb = client.db('author')
    return authorDb
}

export async function getBioDb(): Promise<Db> {
    if (bioDb) {
        return bioDb
    }

    const { client } = await connectToDatabase()
    bioDb = client.db('bio')
    return bioDb
}

export async function getBookDb(): Promise<Db> {
    if (bookDb) {
        return bookDb
    }

    const { client } = await connectToDatabase()
    bookDb = client.db('book')
    return bookDb
}

export async function getContactDb(): Promise<Db> {
    if (contactDb) {
        return contactDb
    }

    const { client } = await connectToDatabase()
    contactDb = client.db('contact')
    return contactDb
}

export async function getProjectDb(): Promise<Db> {
    if (projectDb) {
        return projectDb
    }

    const { client } = await connectToDatabase()
    projectDb = client.db('project')
    return projectDb
}

export async function getSessionDb(): Promise<Db> {
    if (sessionDb) {
        return sessionDb
    }

    const { client } = await connectToDatabase()
    sessionDb = client.db('session')
    return sessionDb
}

export async function getUserDb(): Promise<Db> {
    if (userDb) {
        return userDb
    }

    const { client } = await connectToDatabase()
    userDb = client.db('user')
    return userDb
}

export async function closeDatabase() {
    if (cachedClient) {
        await cachedClient.close()
    }
}
