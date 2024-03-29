import { useState } from 'react'
import { css } from '@emotion/react'
import PacmanLoader from 'react-spinners/PacmanLoader'

const override = css`
	display: block;
	margin: 0 auto;
`

export default function Loading() {
	const [loading, setLoading] = useState(true),
		[color, setColor] = useState<string>('var(--irrigationBlue)')

	return (
		<div className="sweet-loading">
			<PacmanLoader color={color} loading={loading} css={override} size="15" />
		</div>
	)
}
