import { useQuery } from '@apollo/client'
import { GetCategories } from '../../apollo/categories/queries'
import { CategoriesResponse, QueryCategoriesArgs } from '../../generated/graphql'

const Categories = () => {
  const { data, loading, error } = useQuery<CategoriesResponse, QueryCategoriesArgs>(GetCategories)

  return (
    <div>Categories</div>
  )
}

export default Categories