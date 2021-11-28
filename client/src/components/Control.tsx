import { useEffect } from 'react'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'

export default function Control(props: any) {
  useEffect(() => {
    document.title = 'Plant Hub | Control'
  }, [])

  return (
    <div className="control">
      <Card className="card">
        <CardContent>
          <div className="flex-row">
            <div className="flex-col">
              <div className="flex-row pt-5px">
                <span className="title-1">Manuální zavlažování</span>
              </div>
              {/* show only when manual irrigation is active */}
              <div className="flex-row pt-5px">
                <span>Uplynulo času: </span>
              </div>
              <div className="flex-row pt-5px">
                <span>Vody využito: </span>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}
