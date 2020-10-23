import React, { FC, useEffect } from 'react';

import {
  Typography,
  AppBar,
  Toolbar,
  makeStyles,
  createStyles,
  Theme,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  TextField,
  Button,

} from '@material-ui/core';

import {
  Content,
} from '@backstage/core';
import { DefaultApi } from '../../api/apis';
import { EntPatient } from '../../api/models/EntPatient';
import { EntUser } from '../../api/models/EntUser';
import { EntDrug } from '../../api/models/EntDrug';
import Swal from 'sweetalert2'; // alert

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex', //ทำให้อยู่กึ่งกลาง
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    formControl: {
      margin: theme.spacing(1), //ขนาดของบล็อก
      minWidth: 300,
      maxWidth: 300,
      marginTop: '1.5%',
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    title: {
      flexGrow: 1,

    },
    title2: {
      flexGrow: 1,
      marginLeft:60,

    }
  }),
);


interface dispense {
  patient: number;
  drug: number;
  note: string;
  user: number;
}

const Dispense: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();


  const [dispense, setDispense] = React.useState<Partial<dispense>>({});
  const [patients, setPatients] = React.useState<EntPatient[]>([]); //setข้อมูล
  const [drugs, setDrugs] = React.useState<EntDrug[]>([]);
  const [users, setUsers] = React.useState<EntUser[]>([]);


  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,) => {
    const name = event.target.name as keyof typeof Dispense;
    const { value } = event.target;
    setDispense({ ...dispense, [name]: value });
    console.log(dispense);
  };


  const getPatients = async () => {   //ดึงข้อมูล
    const res = await http.listPatient({ limit: 10, offset: 0 });
    setPatients(res);
  };
  const getDrugs = async () => {
    const res = await http.listDrug({ limit: 10, offset: 0 });
    setDrugs(res);
  };
  const getUsers = async () => {
    const res = await http.listUser({ limit: 10, offset: 0 });
    setUsers(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getPatients();
    getDrugs();
    getUsers();
  }, []);

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'center',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  // clear input form
  function clear() {
    setDispense({});
  }
  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/dispenses';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(dispense),
    };

    console.log(dispense); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          //clear();
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        }
      });
  }

  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" className={classes.title}>
            ระบบรายการจ่ายยา
            </Typography>

            <Button variant="contained" href="/" >Logout</Button>

        </Toolbar>
      </AppBar>


      <Content>
        <Typography variant="h4" className={classes.title2}>
          รายการจ่ายยา
            </Typography>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>เลือกชื่อผู้ป่วย</InputLabel>
              <Select
                name="patient"
                label="เลือกชื่อผู้ป่วย"
                value={dispense.patient || ''}
                onChange={handleChange}
              >
                {patients.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>{item.firstname} {item.lastname}</MenuItem>
                  );
                })}
              </Select>
            </FormControl>


            <br></br><br></br>
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>เลือกชื่อยา</InputLabel>
              <Select
                name="drug"
                label="เลือกชื่อยา"
                value={dispense.drug || ''}
                onChange={handleChange}
              >
                {drugs.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>{item.name}</MenuItem>
                  );
                })}
              </Select>
            </FormControl>

            <br></br><br></br>
            <FormControl
              fullWidth
              className={classes.formControl}
              variant="outlined">
              <TextField
                name="note"
                label="หมายเหตุ"
                variant="outlined"
                type="string"
                size="medium"
                value={dispense.note || ''}
                onChange={handleChange}
              />
            </FormControl>

            <br></br><br></br>
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>เลือกชื่อเภสัชกร</InputLabel>
              <Select
                name="user"
                label="เลือกชื่อเภสัชกร"
                value={dispense.user || ''}
                onChange={handleChange}
              >
                {users.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>{item.name}</MenuItem>
                  );
                })}
              </Select>
            </FormControl>


            <br></br><br></br>
            <Button variant="contained" color="primary" onClick={save} style={{ marginRight: 150 }}>SAVE</Button>
            <Button variant="contained" color="secondary" onClick={clear}>CLEAR</Button>



          </form>
        </div>
      </Content>
    </div>

  );
};

export default Dispense;
