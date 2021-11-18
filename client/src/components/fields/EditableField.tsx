import { useState } from 'react'
import { useStyles } from '../../styles/rootStyles'

export default function EditableField({ key, defaultValue }: any) {
  const classes = useStyles()
  const [activeLabel, setActiveLabel] = useState(false)

  return (
    <div className={classes.inputField} onClick={() => setActiveLabel(true)}>
      <input type="text" id="name" className={classes.inputFieldInput} name={key} value={defaultValue} />
    </div>
  )
}
