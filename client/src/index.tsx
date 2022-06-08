import { StrictMode, Suspense } from 'react'
import ReactDOM from 'react-dom'
import App from './App'
import reportWebVitals from './reportWebVitals'
import { ApolloProvider } from '@apollo/client'
import client from './apollo/client'

import './i18n';

ReactDOM.render(
	<StrictMode>
		<ApolloProvider client={client}>
			<Suspense fallback="loading">
				<App />
			</Suspense>
		</ApolloProvider>
	</StrictMode>,
	document.getElementById('root')
)

reportWebVitals()
