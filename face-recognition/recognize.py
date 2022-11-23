import face_recognition
import imutils
import pickle
import time
import cv2
import os
 
faceCascade = cv2.CascadeClassifier("./face-recognition/haarcascade_frontalface_alt2.xml")

data = pickle.loads(open('./face-recognition/face_enc', "rb").read())

#print("Streaming started")
video_capture = cv2.VideoCapture(0)
t_end = time.time() + 5
while time.time() < t_end:
    ret, frame = video_capture.read()
    gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
    faces = faceCascade.detectMultiScale(gray,
                                         scaleFactor=1.1,
                                         minNeighbors=5,
                                         minSize=(60, 60),
                                         flags=cv2.CASCADE_SCALE_IMAGE)
 
    rgb = cv2.cvtColor(frame, cv2.COLOR_BGR2RGB)
    # the facial embeddings for face in input
    encodings = face_recognition.face_encodings(rgb)
    names = []
    # loop over the facial embeddings incase
    # we have multiple embeddings for multiple fcaes
    for encoding in encodings:
       #Compare encodings with encodings in data["encodings"]
       #Matches contain array with boolean values and True for the embeddings it matches closely
       #and False for rest
        matches = face_recognition.compare_faces(data["encodings"],
         encoding)
        #set name =inknown if no encoding matches
        name = "Unknown"
        # check to see if we have found a match
        if True in matches:
            #Find positions at which we get True and store them
            matchedIdxs = [i for (i, b) in enumerate(matches) if b]
            counts = {}
            # loop over the matched indexes and maintain a count for
            # each recognized face face
            for i in matchedIdxs:
                #Check the names at respective indexes we stored in matchedIdxs
                name = data["names"][i]
                #increase count for the name we got
                counts[name] = counts.get(name, 0) + 1
            #set name which has highest count
            name = max(counts, key=counts.get)
 
 
        # update the list of names
        #names.append(name)
        # loop over the recognized faces
        # for ((x, y, w, h), name) in zip(faces, names):
        #     # rescale the face coordinates
        #     # draw the predicted face name on the image
        #     cv2.rectangle(frame, (x, y), (x + w, y + h), (0, 255, 0), 2)
        #     cv2.putText(frame, name, (x, y), cv2.FONT_HERSHEY_SIMPLEX,
        #      0.75, (0, 0, 0), 2)
    cv2.imshow("Frame", frame)
    if cv2.waitKey(1) & 0xFF == ord('q'):
        break
print(name)
video_capture.release()
cv2.destroyAllWindows()