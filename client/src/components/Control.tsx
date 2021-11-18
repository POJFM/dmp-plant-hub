import { useEffect } from 'react'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import { useStyles } from './../styles/rootStyles'
import { useControlStyles } from './../styles/control'

export default function Control(props: any) {
  const controlClasses = useControlStyles()
  const classes = useStyles()
  useEffect(() => {
    document.title = 'Plant Hub | Control'
  }, [])

  return (
    <div className="col control">
      <Card className={classes.card}>
        <CardContent>
          <div className="row">
            <div className="col">
              <div className={`row ${classes.cardRow}`}>
                <span>Manual irrigation</span>
              </div>
              {/* show only when manual irrigation is active */}
              <div className={`row ${classes.cardRow}`}>
                <span>Time passed: </span>
              </div>
              <div className={`row ${classes.cardRow}`}>
                <span>Water overdrown: </span>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}
