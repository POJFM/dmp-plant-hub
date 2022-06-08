import { useState, useEffect } from 'react'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import { makeStyles } from "@material-ui/core/styles";

export const useStyles = makeStyles({
  bg: {
    background: 'var(--white)'
  },
  bgDark: {
    background: 'var(--cardGreen)'
  }
});

export const MuiCard = ({ children }: any) => {
  const classes = useStyles(),
    [theme, setTheme] = useState(localStorage.getItem('theme'))

  useEffect(() => {
    const themeState = localStorage.getItem('theme')
    themeState && setTheme(themeState)
  })

  return (
    <Card className="card" classes={{ root: theme === 'dark' ? classes.bgDark : classes.bg }}>
      <CardContent>{children}</CardContent>
    </Card>
  )
}