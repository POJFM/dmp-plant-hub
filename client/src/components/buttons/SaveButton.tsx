import { useState } from 'react'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import { useStyles } from '../../styles/rootStyles'

export default function SaveButton({ ...props }: any) {
  const classes = useStyles()
  const [buttonHover, setButtonHover] = useState(false)

  return (
    <div className={classes.buttonWrapper}>
      <Card
        className={`${classes.button} ${buttonHover && classes.buttonHover}`}
        onMouseOver={() => setButtonHover(true)}
        onMouseOut={() => setButtonHover(false)}
      >
        <CardContent>
          <span className={classes.buttonText}>Uložit</span>
        </CardContent>
      </Card>
    </div>
  )
}
