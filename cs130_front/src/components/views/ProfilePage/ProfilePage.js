import React, {useState} from 'react';
import './styles.css';
import Requests from '../../Requests/Requests';
import Infos from '../../Infos/Infos';
import SimpleForm from '../../Infos/SimpleForm';
import ClassList from '../../CurrentClasses/ClassList';
import MyListings from './MyListings';
import SchedulerPage from '../SchedulerPage/SchedulerPage';
import CourseAdder from './CourseAdder';
import CourseView from './CourseView';
import * as Colors from '../../../constants/Colors';
import Text from '../../Text/Text';
import Button from '../../Button/Button';
import { forEach } from 'lodash';

function ProfilePage(props) {
  const invs = [{name: "Al Squad", id: 223, types: "invitation"}, {name: "Calc Gang", id: 223, types: "invitation"}]
  const contactInformation = [
    {name: "Email", value: "Edgar@gmail.com"}, 
    {name: "Facebook", value: "facebook.com/Edgar"},
    {name: "Discord", value: "Edgar#1234"}
  ];
  const additionalInformation = [
    {name: "Name", value: "Edgar Garcia"},
    {name: "School name", value: "UCLA"}, 
    {name: "Timezone", value: "PST"},
    {name: "Biography", value: "I enjoy graphs and cookies"},
  ];
  var availability1 = [1,1,1,1,1,1];
  var availability2 = new Array(330).fill(0);
  var fullAvailability = availability1.concat(availability2);
  const l = [
    {id:123, courseName: "Algebra I", 
    content: "HMU PLS I really want to talk about derivatves with people from all over the country. I am from Wisconsin. I am very cool."},
    {id:124, courseName: "Calculus I", content: "HMU PLS i love calculus"}
  ];

  const classesList = [
    {
      institution: "College",
      categories: [
        {
          category: "Mathematics",
          classes: [
            {name: "Discrete Mathematics", classId: 123, keywords:["graphs", "acyclis", "paths"]},
            {name: "Integral Calculus", classId: 124, keywords:["double integral", "radial", "Greene thereom"]},
            {name: "Derivative Calculus", classId: 125, keywords:["area of a curve", "limits", "partial derivative"]}
          ]
        },
        {
          category: "Science",
          classes: [
            {name: "Organic Chemistry", classId: 126, keywords:["graphs", "acyclis", "paths"]},
          ]
        },
        {
          category: "Psychology",
          classes: []
        },
      ]
    },
    {
      institution: "High School",
      categories: [
        {
          category: "Mathematics",
          classes: [
            {name: "Algebra I", classId: 1231, keywords:["find x", "zeros", "idk", "derivatives", "single integral", "partial derivative", "derivatives", "single integral", "partial derivative"]},
            {name: "Algebra II", classId: 1241, keywords:["idk", "whats in", "this class"]},
            {name: "AP Calculus AB", classId: 1251, keywords:["derivatives", "single integral", "partial derivative"]}
          ]
        }
      ]
    },
  ];

  const myCourses2= [{name: "Algebra I", classId: 1231, keywords:["find x", "zeros", "idk", "derivatives", "single integral", "partial derivative", "derivatives", "single integral", "partial derivative"], groups:[{name: "Al Squad", id: 123}]},
  {name: "Calculus I", classId: 1231, keywords:["find x", "zeros", "idk", "derivatives", "single integral", "partial derivative", "derivatives", "single integral", "partial derivative"]}
  ];

  const [invitations, setInvitations] = useState(invs);
  const [contactInfo, setContactInfo] = useState(contactInformation);
  const [additionalInfo, setAdditionalInfo] = useState(additionalInformation);
  const [listings, setListings] = useState(l);
  const [classes, setClasses] = useState(classesList);
  const [myCourses, setMyCourses] = useState(myCourses2);
  const [mainPanelState, setMainPanel] = useState('CourseAdder');
  const userId = props.match.params.id;
  const myId = localStorage.getItem('userId');
  // check clickable by checking user's id with path's id

  function getProfile() {
    setInvitations(invs);
    setContactInfo(contactInfo);
    setAdditionalInfo(additionalInfo);
    setListings(l);
    setMyCourses(myCourses2)
  }

  function getClassesList() {
    setClasses(classesList);
  }

  function renderMainPanel() {
    switch (mainPanelState) {
      case 'Contact Information':
        return <SimpleForm options={contactInfo} saveInfoClicked={saveInfoClicked}/>
      case 'Additional Information': 
        return <SimpleForm options={additionalInfo} saveInfoClicked={saveInfoClicked}/>
      case 'CourseAdder': 
        return <CourseAdder courses={classes} addCourse={addCourse}/>
      default: //Course view for selected course
        const course = myCourses.find(element => element.name == mainPanelState.name)
        return <CourseView item={course} removeCourse={removeCourse}/>
    }
  }

  function coursesOnly() {
    var newCourses = [...myCourses2];
    newCourses.forEach( element => {
      delete element["groups"]
    })
    return newCourses;
  }

  function titleClicked(title) {
    //set main content to be for title
    setMainPanel(title);
  }

  function groupClicked(item) {
    //navigate to group page
    console.log(item);
  }

  function saveInfoClicked(first, second, third) {
    //post to endpoint updating info
  }

  function handleInvitation(invitation) {
    //update invitation with accept/decline
    // and remove from list then fetch again to update ui
  }

  function editListing(content, listing) {
    console.log(content, listing)
    //edit listing endpoint
  }
  
  function addCourse(course) {
    console.log(course);
  }

  function removeCourse(course) {

  }

  function viewOnlyProfile() {
    return (
      <div className="view-panel">
        <div className="column"> 
          <div className="text-container">
            <Text color="black" size="32px" weight="800"> 
            {additionalInfo[0].value} Profile
            </Text>
            <ClassList classList={userId == myId ? coursesOnly() : myCourses} titleClicked={groupClicked} clickable={false}/>
            <Infos title="Contact Information" options={contactInformation} titleClicked={titleClicked} clickable={false}/>
            <Infos title="Additional Information" options={additionalInformation} titleClicked={titleClicked} clickable={false}/>
          </div>
        </div>
        <div className="column">
          <SchedulerPage passedSelections={fullAvailability}/>
        </div>
      </div>
    );

  }

  function myProfile() {
    return (
      <div className="panel">
  
        <div className="column"> 
          <div className="text-container">
            <Text color="black" size="24px" weight="800"> 
              Hi Student, edit your class and other information to start studying with others!
            </Text>
          </div>
          <MyListings items={listings} editListing={editListing}/>
          <Requests title="Invitations" items={invitations} handleResponse={handleInvitation}/>
        </div>
        <div className="column"> 
          {renderMainPanel()}
        </div>
        <div className="column">
          <div className="group-with-margin-bottom">
            <ClassList classList={userId == myId ? coursesOnly() : myCourses} titleClicked={titleClicked} clickable={true}/>
            <Infos title="Contact Information" options={contactInformation} titleClicked={titleClicked} clickable={true}/>
            <Infos title="Additional Information" options={additionalInformation} titleClicked={titleClicked} clickable={true}/>
          </div>
          <div className="group-with-margin-bottom">
            <Button 
            textColor={Colors.White}
            textSize="28px"
            width="280px"
            height="70px"
            textWeight="800" 
            color={Colors.Blue}
            onClick={() => setMainPanel('CourseAdder')}
            >
              Add Courses
            </Button>
          </div>
          <Button 
          textColor={Colors.White}
          textSize="28px"
          width="280px"
          height="70px"
          textWeight="800" 
          color={Colors.Blue}
          onClick={() => props.history.push({
            pathname: `/profile/${userId}/scheduler`,
            state: { availability: fullAvailability }
            })}
          >
            Set Availability
          </Button>
        </div>
      </div>
    );
  }

  // will have conditional logic based on self profile vs different.. maybe a diff endpoint?
  return userId == myId ? myProfile() : viewOnlyProfile();
}

export default ProfilePage;
