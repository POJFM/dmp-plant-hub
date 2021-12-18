import CancelIcon from '@material-ui/icons/Cancel'

export default function Warning() {
	return (
    <div className="text-4xl flex items-center duration-500" style={{color: 'var(--warningRed)'}}>
      <CancelIcon/>
    </div>
  )
}
