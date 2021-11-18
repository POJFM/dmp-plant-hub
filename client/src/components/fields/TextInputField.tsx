import { useState } from 'react'
import { useStyles } from '../../styles/rootStyles'

export default function TextInputField({ key, name, defaultValue }: any) {
  const classes = useStyles()
  const [activeLabel, setActiveLabel] = useState(false)

  return (
    <div className={classes.inputField} onClick={() => setActiveLabel(true)}>
      <label htmlFor="name" className={classes.inputFieldLabel}>
        {name}
      </label>
      <input type="text" id="name" className={classes.inputFieldInput} name={key} value={defaultValue} />
    </div>
  )
}
