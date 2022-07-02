import { gql } from '@apollo/client';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  DateTime: any;
};

export type Article = {
  __typename?: 'Article';
  categoryId?: Maybe<Scalars['Int']>;
  createdAt?: Maybe<Scalars['DateTime']>;
  html?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['Int']>;
  image?: Maybe<Scalars['String']>;
  status?: Maybe<Status>;
  text?: Maybe<Scalars['String']>;
  title?: Maybe<Scalars['String']>;
  updatedAt?: Maybe<Scalars['DateTime']>;
  url?: Maybe<Scalars['String']>;
  userId?: Maybe<Scalars['Int']>;
};

export type CategoriesResponse = {
  __typename?: 'CategoriesResponse';
  items?: Maybe<Array<Maybe<Category>>>;
  total?: Maybe<Scalars['Int']>;
};

export type Category = {
  __typename?: 'Category';
  createdAt?: Maybe<Scalars['DateTime']>;
  id?: Maybe<Scalars['Int']>;
  title?: Maybe<Scalars['String']>;
  userId?: Maybe<Scalars['Int']>;
};

export type FilterSchemaInput = {
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  sortBy?: InputMaybe<Scalars['String']>;
  sortOrder?: InputMaybe<Scalars['String']>;
};

export type Mutation = {
  __typename?: 'Mutation';
  /** Add a new article */
  addArticle?: Maybe<Article>;
  /** Add a new category */
  addCategory?: Maybe<Category>;
  /** Log in a user */
  loginUser?: Maybe<User>;
  /** Sign up a user */
  signUpUser?: Maybe<User>;
};


export type MutationAddArticleArgs = {
  categoryId: Scalars['Int'];
  url: Scalars['String'];
};


export type MutationAddCategoryArgs = {
  title: Scalars['String'];
};


export type MutationLoginUserArgs = {
  password: Scalars['String'];
  usernameOrEmail: Scalars['String'];
};


export type MutationSignUpUserArgs = {
  email: Scalars['String'];
  password: Scalars['String'];
  username: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  article?: Maybe<Article>;
  categories?: Maybe<CategoriesResponse>;
  category?: Maybe<Category>;
  /** Get current user */
  me?: Maybe<User>;
};


export type QueryArticleArgs = {
  id: Scalars['Int'];
};


export type QueryCategoriesArgs = {
  filter?: InputMaybe<FilterSchemaInput>;
};


export type QueryCategoryArgs = {
  id: Scalars['ID'];
};

export type User = {
  __typename?: 'User';
  createdAt?: Maybe<Scalars['DateTime']>;
  email?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['Int']>;
  username?: Maybe<Scalars['String']>;
};

/** Article status */
export enum Status {
  /** article is read */
  Read = 'read',
  /** article is not read yet */
  Unread = 'unread'
}
