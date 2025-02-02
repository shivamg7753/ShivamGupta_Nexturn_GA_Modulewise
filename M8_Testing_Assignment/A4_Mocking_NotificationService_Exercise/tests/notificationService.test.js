const sendNotification = require('../src/notificationService');

describe('sendNotification', () => {
  let mockNotificationService;

  beforeEach(() => {
    // Create a mock for the notification service before each test
    mockNotificationService = {
      send: jest.fn(),
    };
  });

  test('should return "Notification Sent" when send is successful', () => {
    // Mock the send method to simulate a successful notification
    mockNotificationService.send.mockReturnValue(true);

    const message = 'Hello, World!';
    const result = sendNotification(mockNotificationService, message);

    expect(result).toBe('Notification Sent');
    expect(mockNotificationService.send).toHaveBeenCalledWith(message);
  });

  test('should return "Failed to Send" when send fails', () => {
    // Mock the send method to simulate a failed notification
    mockNotificationService.send.mockReturnValue(false);

    const message = 'Hello, World!';
    const result = sendNotification(mockNotificationService, message);

    expect(result).toBe('Failed to Send');
    expect(mockNotificationService.send).toHaveBeenCalledWith(message);
  });
});
