import React, { FC, useEffect, useReducer } from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import { makeStyles } from '@material-ui/core/styles';
import { EntUser, EntUserFromJSON } from '../../api/models/EntUser';
import {
  // Content,
  Header,
  Page,
  pageTheme,
  // ContentHeader,
} from '@backstage/core';
import { DefaultApi } from '../../api';
import User from '../Users';
const HeaderCustom = {
  minHeight: '120px',
};
const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  form: {
    width: '100%',
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));
interface State {
  email: string;
  password: string;
  isButtonDisabled: boolean
  helperText: string
  isError: boolean
}
const initialState: State = {
  email: '',
  password: '',
  isButtonDisabled: true,
  helperText: '',
  isError: false,
}
type Action = { type: 'setEmail', payload: string }
  | { type: 'setPassword', payload: string }
  | { type: 'setIsButtonDisabled', payload: boolean }
  | { type: 'loginSuccess', payload: string }
  | { type: 'loginFailed', payload: string }
  | { type: 'setIsError', payload: boolean };
const reducer = (state: State, action: Action): State => {
  switch (action.type) {
    case 'setEmail':
      return {
        ...state,
        email: action.payload
      };
    case 'setPassword':
      return {
        ...state,
        password: action.payload
      };
    case 'setIsButtonDisabled':
      return {
        ...state,
        isButtonDisabled: action.payload
      };
    case 'loginSuccess':
      return {
        ...state,
        helperText: action.payload,
        isError: false
      };
    case 'loginFailed':
      return {
        ...state,
        helperText: action.payload,
        isError: true
      };
    case 'setIsError':
      return {
        ...state,
        isError: action.payload
      };
  }
}

const Login = () => {
  const classes = useStyles();
  const [state, dispatch] = useReducer(reducer, initialState);
  const [user, ] = React.useState<EntUser[]>([]);
  const [password, getPassword] = React.useState<EntUser[]>([]);

  
  useEffect(() => {
    if (state.email.trim() && state.password.trim()) {
      dispatch({
        type: 'setIsButtonDisabled',
        payload: false
      });
    } else {
      dispatch({
        type: 'setIsButtonDisabled',
        payload: true
      });
    }
  }, [state.email, state.password]);
  const handleLogin = () => {
    if ((state.email === 'w@gmail.com' && state.password === '12345') ||
        (state.email === 'raweewan@gmail.com' && state.password === '12345')) {
      dispatch({
        type: 'loginSuccess',
        payload: 'Login Successfully'
      });
    } else {
      dispatch({
        type: 'loginFailed',
        payload: 'Incorrect username or password'
      });
    }
  };

 
  const handleKeyPress = (event: React.KeyboardEvent) => {
    if (event.keyCode === 13 || event.which === 13) {
      state.isButtonDisabled || handleLogin();
    }
  };
  const handleEmailChange: React.ChangeEventHandler<HTMLInputElement> =
    (event) => {
      dispatch({
        type: 'setEmail',
        payload: event.target.value
      });
    };
  const handlePasswordChange: React.ChangeEventHandler<HTMLInputElement> =
    (event) => {
      dispatch({
        type: 'setPassword',
        payload: event.target.value
      });
    }
  return (
    <div className={classes.paper}>
      <Page theme={pageTheme.website}>
        <Header style={HeaderCustom} title={`Medicine Room System`}
          subtitle="กรุณาบันทึกข้อมูลก่อนเข้าสู่ระบบ">
        </Header>
        <form className={classes.submit} noValidate>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            autoComplete="email"
            autoFocus
            onChange={handleEmailChange}
            onKeyPress={handleKeyPress}

          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
            helperText={state.helperText}
            onChange={handlePasswordChange}
            onKeyPress={handleKeyPress}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            onClick={handleLogin}
            disabled={state.isButtonDisabled}
            href="/user"
            >
            Sign In
          </Button>
          <Grid container>
            <Grid item>
              <Link href="/user" variant="body2">
                {"Don't have an account? Sign Up"}
              </Link>
            </Grid>
          </Grid>
        </form>
      </Page>
    </div>
  );
};
export default Login;