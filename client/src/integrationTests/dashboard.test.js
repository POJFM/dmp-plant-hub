import { render, screen, fireEvent } from '@testing-library/react'
import '@testing-library/jest-dom'
import Dashboard from './../components/Dashboard'
import { ApolloProvider } from '@apollo/client'
import client from './../apollo/client'

test('Dashboard logic', async () => {
	render(
		<ApolloProvider client={client}>
			<Dashboard />
		</ApolloProvider>
	)

	global.XMLHttpRequest = undefined
})
