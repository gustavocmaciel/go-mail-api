-- Mock data for the 'users' table
INSERT INTO users (id, email, first_name, last_name) VALUES
  ('4a3a51c7-0e9f-4b6b-8a36-87bb5be5c054', 'user1@example.com', 'John', 'Doe'),
  ('fc3b3df4-58d1-4e01-a0f7-21b1cda827f0', 'user2@example.com', 'Alice', 'Smith'),
  ('9fe4b2b0-13b7-4e8d-a0b9-8c1a4078ff2b', 'user3@example.com', 'Bob', 'Johnson');

-- Mock data for the 'emails' table
INSERT INTO emails (id, sender, recipients, subject, body, timestamp, email_read, archived) VALUES
  ('3a6f9df1-25f8-4922-928d-056f76954284', 'user1@example.com', ARRAY['user2@example.com', 'user3@example.com'], 'Meeting Invitation', 'Please join us for the meeting.', '2023-01-01 10:00:00', true, false),
  ('8a89df90-7a98-4ab1-a27d-b98f8e102a3d', 'user2@example.com', ARRAY['user1@example.com'], 'Project Update', 'Here is the latest update on the project.', '2023-01-02 15:30:00', false, false),
  ('c7e1106a-1f4e-4c23-b1b6-f21cb0f068ec', 'user3@example.com', ARRAY['user1@example.com', 'user2@example.com'], 'Important Announcement', 'Read this important announcement.', '2023-01-03 12:45:00', false, true);
