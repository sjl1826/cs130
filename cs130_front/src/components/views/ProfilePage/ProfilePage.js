import React, {useState} from 'react';
import './styles.css';
import Requests from '../../Requests/Requests';
import Infos from '../../Infos/Infos';
import SimpleForm from '../../Infos/SimpleForm';
import MyListings from './MyListings';
import CourseAdder from './CourseAdder';
import CourseView from './CourseView';
import * as Colors from '../../../constants/Colors';
import Text from '../../Text/Text';
import Button from '../../Button/Button';

function ProfilePage(props) {
  const invs = [{name: "CA Squad", id: 223, types: "invitation"}, {name: "Alexander Hamilton-Tuff", id: 223, types: "invitation"}]
  const contactInformation = [
    {name: "Email", value: "shirly@gmail.com"}, 
    {name: "Facebook", value: "facebook.com/shirly"},
    {name: "Discord", value: "shirly#1234"}
  ];
  const additionalInformation = [
    {name: "School name", value: "UCLA"}, 
    {name: "Timezone", value: "PST"},
    {name: "Biography", value: "I enjoy graphs and cookies"}
  ];
  var availability1 = [1,1,1,1,1,1];
  var availability2 = new Array(330).fill(0);
  var fullAvailability = availability1.concat(availability2);
  const l = [
    {id:123, courseName: "Discrete Mathematics", 
    content: "HMU PLS I really want to talk about bipartite graphs with people from all over the country. I am from Wisconsin. I am very cool."},
    {id:124, courseName: "Human Anatomy", content: "HMU PLS i love anatomy"}
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

  const myCourses= [{name: "Algebra I", classId: 1231, keywords:["find x", "zeros", "idk", "derivatives", "single integral", "partial derivative", "derivatives", "single integral", "partial derivative"]}];

  const [invitations, setInvitations] = useState(invs);
  const [contactInfo, setContactInfo] = useState(contactInformation);
  const [additionalInfo, setAdditionalInfo] = useState(additionalInformation);
  const [listings, setListings] = useState(l);
  const [classes, setClasses] = useState(classesList);
  const [mainPanelState, setMainPanel] = useState('Classes');
  const userId = props.match.params.id;
  // check clickable by checking user's id with path's id


  function getProfile() {
    setInvitations(invs);
    setContactInfo(contactInfo);
    setAdditionalInfo(additionalInfo);
    setListings(l);
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
        return <CourseView item={myCourses[0]} removeCourse={removeCourse}/>
    }
  }

  function titleClicked(title) {
    //set main content to be for title
    setMainPanel(title);
  }

  function saveInfoClicked(first, second, third) {
    //post to endpoint updating info
  }

  function handleInvitation(invitation) {
    //update invitation with accept/decline
    // and remove from list then fetch again to update ui
  }

  function editListing(listing) {
    //edit listing endpoint
  }

  function addCourse(course) {

  }

  function removeCourse(course) {

  }

  // will have conditional logic based on self profile vs different.. maybe a diff endpoint?
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
        <div className="group-with-margin">
          <Infos title="Contact Information" options={contactInformation} titleClicked={titleClicked} clickable={true}/>
          <Infos title="Additional Information" options={additionalInformation} titleClicked={titleClicked} clickable={true}/>
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

export default ProfilePage;