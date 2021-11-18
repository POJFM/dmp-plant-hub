import { useEffect } from 'react'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import { useStyles } from './../styles/rootStyles'

export default function History(props: any):any {
  const classes = useStyles()
  useEffect( () => {
    document.title = 'Plant Hub | History'
  }, [])

  return (
    <div className='row history'>
      <div className='col'>
        <Card>
          <CardContent>

          </CardContent>
        </Card>
      </div>
    </div>
  )
}