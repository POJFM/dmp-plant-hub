import { render, screen, fireEvent } from '@testing-library/react'
import '@testing-library/jest-dom'
import Control from './../components/Control'
import { ApolloProvider } from '@apollo/client'
import client from './../apollo/client'

test('Control logic', async () => {
	render(
		<ApolloProvider client={client}>
			<Control />
		</ApolloProvider>
	)

	global.XMLHttpRequest = undefined
})
