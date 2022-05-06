import { ApolloClient, InMemoryCache, HttpLink } from '@apollo/client'
import fetch from 'cross-fetch'

const client = new ApolloClient({
	link: new HttpLink({ uri: `${process.env.REACT_APP_GO_API_URL}/query`, fetch }),
	cache: new InMemoryCache(),
})

export default client
