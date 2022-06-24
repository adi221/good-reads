import Snackbar from '@mui/material/Snackbar';
import MuiAlert from '@mui/material/Alert';
import { useMessage } from '../../contexts/MessageContext';


const AlertCenter = () => {
  const { message, showMessage } = useMessage()

  const handleClose = () => {
    showMessage({ text: '' });
  };

  return (
    <Snackbar onClose={handleClose} open={!!message.text} autoHideDuration={message.ttl}>
      <MuiAlert onClose={handleClose} elevation={6} variant="filled" severity={message.variant} sx={{ width: '100%' }}>
        {message.text}
      </MuiAlert>
    </Snackbar>
  )
}

export default AlertCenter