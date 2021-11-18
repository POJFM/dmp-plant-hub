import { useState } from 'react'
import { useStyles } from '../../styles/rootStyles'


export default function EnumEditableField(key: any, values: any) {
  const classes = useStyles()

  return (
    // {values.map(value: any) => {
    //   return(

    //   )
    // }}
    <div className={classes.inputField}>
    <input type="text" id="theme" className={classes.inputFieldInput} name={key} /* value={defaultValue} */ />
  </div>
  )
}
