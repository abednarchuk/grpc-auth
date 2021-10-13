import { useEffect } from 'react'
import * as grpcWeb from 'grpc-web';
import { SignupRequest } from './authpb/auth_pb'
import { SignupServiceClient } from './authpb/AuthServiceClientPb'
import './App.css';

function App() {
  let signupServiceClient = new SignupServiceClient('http://localhost:8000')
  const request = new SignupRequest()
  request.setEmail("test@test.com")
  request.setPassword("secret")

  useEffect(() => {
    signupServiceClient.signUp(request, {}, (err, res) => {
      try {
        console.log(err)
        console.log(res)
      } catch (err) { }
    })
  }, [])

  return (
    <div>
      test
    </div>
  );
}

export default App;
