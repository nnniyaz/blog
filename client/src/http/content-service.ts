import { $api, ApiRoutes } from '@/http/index'
import { ArticlesMany } from '@domain/article/article'
import { Author } from '@domain/author/author'
import { QueryParams } from '@domain/base/query-params'
import { ErrorResponse, SuccessResponse } from '@domain/base/response'
import { Bio } from '@domain/bio/bio'
import { BooksMany } from '@domain/book/book'
import { Contact, ContactsMany } from '@domain/contact/contact'
import { ProjectsMany } from '@domain/project/project'
import { AxiosResponse } from 'axios'

export class ContentServiceHttp {
    static async getArticles(
        params: QueryParams,
    ): Promise<AxiosResponse<SuccessResponse<ArticlesMany> | ErrorResponse>> {
        return $api.get(ApiRoutes.ARTICLE, { params })
    }

    static async getAuthor(): Promise<
        AxiosResponse<SuccessResponse<Author> | ErrorResponse>
    > {
        return $api.get(ApiRoutes.AUTHOR)
    }

    static async getBio(): Promise<
        AxiosResponse<SuccessResponse<Bio> | ErrorResponse>
    > {
        return $api.get(ApiRoutes.BIO)
    }

    static async getBooks(
        params: QueryParams,
    ): Promise<AxiosResponse<SuccessResponse<BooksMany> | ErrorResponse>> {
        return $api.get(ApiRoutes.BOOK, { params })
    }

    static async getContacts(): Promise<
        AxiosResponse<SuccessResponse<ContactsMany> | ErrorResponse>
    > {
        return $api.get(ApiRoutes.CONTACT)
    }

    static async getProjects(
        params: QueryParams,
    ): Promise<AxiosResponse<SuccessResponse<ProjectsMany> | ErrorResponse>> {
        return $api.get(ApiRoutes.PROJECT, { params })
    }
}

export class ContentService {
    static async getArticles(
        params: QueryParams,
    ): Promise<ArticlesMany | null> {
        try {
            const response = await ContentServiceHttp.getArticles(params)
            if (response.data.success) {
                return response.data.data
            }
            return null
        } catch (error) {
            console.log(error)
            return null
        }
    }

    static async getAuthor(): Promise<Author | null> {
        try {
            const response = await ContentServiceHttp.getAuthor()
            if (response.data.success) {
                return response.data.data
            }
            return null
        } catch (error) {
            console.log(error)
            return null
        }
    }

    static async getBio(): Promise<Bio | null> {
        try {
            const response = await ContentServiceHttp.getBio()
            if (response.data.success) {
                return response.data.data
            }
            return null
        } catch (error) {
            console.log(error)
            return null
        }
    }

    static async getBooks(params: QueryParams): Promise<BooksMany | null> {
        try {
            const response = await ContentServiceHttp.getBooks(params)
            if (response.data.success) {
                return response.data.data
            }
            return null
        } catch (error) {
            console.log(error)
            return null
        }
    }

    static async getContacts(): Promise<Contact[] | null> {
        try {
            const response = await ContentServiceHttp.getContacts()
            if (response.data.success) {
                return response.data.data.contacts
            }
            return null
        } catch (error) {
            console.log(error)
            return null
        }
    }

    static async getProjects(
        params: QueryParams,
    ): Promise<ProjectsMany | null> {
        try {
            const response = await ContentServiceHttp.getProjects(params)
            if (response.data.success) {
                return response.data.data
            }
            return null
        } catch (error) {
            console.log(error)
            return null
        }
    }
}
