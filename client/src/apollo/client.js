import { ApolloClient, InMemoryCache } from '@apollo/client'

const client = new ApolloClient({
	uri: `http://4.2.0.225:5000/query`,
	cache: new InMemoryCache(),
})

export default client
