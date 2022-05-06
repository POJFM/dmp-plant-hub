import { render, screen, fireEvent } from '@testing-library/react'
import '@testing-library/jest-dom'
import Settings from './../components/Settings'
import { ApolloProvider } from '@apollo/client'
import client from './../apollo/client'

test('Settings logic', async () => {
	render(
		<ApolloProvider client={client}>
			<Settings />
		</ApolloProvider>
	)

	global.XMLHttpRequest = undefined
})
