import { useState, useCallback } from 'react'
import { GoogleMap, Marker, useJsApiLoader } from '@react-google-maps/api'

interface coords {
	lat: number
	lon: number
}

export default function Map({ lat, lon }: coords) {
	const { isLoaded } = useJsApiLoader({
		id: 'google-map-script',
		googleMapsApiKey: process.env.REACT_APP_GOOGLE_API_KEY || '',
	})

	const [map, setMap] = useState(null)
	const [zoomUpdate, setZoomUpdate] = useState(false)

	const onLoad = useCallback((map) => {
		const bounds = new window.google.maps.LatLngBounds()
		map.fitBounds(bounds)
		setMap(map)
	}, [])

	const onUnmount = useCallback((map) => {
		setMap(null)
	}, [])

	return isLoaded ? (
		<GoogleMap
			mapContainerStyle={{
				width: '300px',
				height: '600px',
			}}
			center={{ lat: lat, lng: lon }}
			onLoad={onLoad}
			onUnmount={onUnmount}
		>
			<Marker position={{ lat: lat, lng: lon }} />
		</GoogleMap>
	) : (
		<></>
	)
}
