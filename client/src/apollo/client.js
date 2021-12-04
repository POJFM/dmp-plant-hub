import { ApolloClient, InMemoryCache } from '@apollo/client'
//import { resolvers } from '../graphql/resolvers'

const client = new ApolloClient({
	uri: 'https://serene-carlsbad-caverns-55233.herokuapp.com/graphql',
	cache: new InMemoryCache(),
	//resolvers,
	// fetchOptions: {
	// 	mode: 'no-cors',
	// },
})

export default client